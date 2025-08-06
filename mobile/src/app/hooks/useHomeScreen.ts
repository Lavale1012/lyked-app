import { useState, useEffect, useCallback, useMemo } from 'react';
import { Upload, HomeScreenState, UserStats } from '../types/HomeScreen.types';
import { useUser } from '../context/UserContext';

const API_KEY = process.env.EXPO_PUBLIC_LINKPREVIEW_API_KEY;
const API_BASE_URL = process.env.EXPO_PUBLIC_API_BASE_URL || 'http://localhost:8080';
const DEFAULT_IMAGE =
  'https://media.istockphoto.com/id/1147544807/vector/thumbnail-image-vector-graphic.jpg?s=612x612&w=0&k=20&c=rnCKVbdxqkjlcs3xH87-9gocETqpspHFXu5dIGB4wuM=';

interface UseHomeScreenReturn {
  state: HomeScreenState;
  searchQuery: string;
  filteredUploads: Upload[];
  greeting: string;
  setSearchQuery: (query: string) => void;
  fetchUploads: (showRefreshControl?: boolean) => Promise<void>;
  handleRefresh: () => void;
}

export const useHomeScreen = (): UseHomeScreenReturn => {
  const { userId } = useUser();
  
  const [state, setState] = useState<HomeScreenState>({
    uploads: [],
    recentUploads: [],
    isLoading: true,
    isRefreshing: false,
    hasError: false,
    errorMessage: '',
    userStats: {
      totalSavedItems: 0,
      recentActivityCount: 0,
    },
  });
  
  const [searchQuery, setSearchQuery] = useState('');

  // Generate greeting based on time of day
  const greeting = useMemo((): string => {
    const hour = new Date().getHours();
    
    if (hour < 6) return 'Working Late? 🌙';
    if (hour < 12) return 'Good Morning! ☀️';
    if (hour < 18) return 'Good Afternoon! 🌤️';
    return 'Good Evening! 🌅';
  }, []);

  // Filter uploads based on search query
  const filteredUploads = useMemo(() => {
    if (!searchQuery.trim()) return state.uploads;
    
    const query = searchQuery.toLowerCase();
    return state.uploads.filter(upload =>
      upload.title.toLowerCase().includes(query) ||
      upload.description.toLowerCase().includes(query) ||
      upload.tags?.some(tag => tag.toLowerCase().includes(query)) ||
      upload.category?.toLowerCase().includes(query)
    );
  }, [state.uploads, searchQuery]);

  // Fetch uploads with enhanced error handling and caching
  const fetchUploads = useCallback(async (showRefreshControl = false) => {
    if (!API_KEY) {
      console.error('LinkPreview API key is missing');
      setState(prev => ({
        ...prev,
        isLoading: false,
        hasError: true,
        errorMessage: 'API configuration error',
      }));
      return;
    }

    setState(prev => ({
      ...prev,
      isRefreshing: showRefreshControl,
      isLoading: !showRefreshControl ? true : prev.isLoading,
      hasError: false,
      errorMessage: '',
    }));

    try {
      const response = await fetch(`${API_BASE_URL}/uploads/all?user_id=${userId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Cache-Control': 'no-cache',
          'Origin': API_BASE_URL,
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      const uploadsData = data.uploads || [];

      // Batch process uploads for better performance
      const processUpload = async (upload: Upload): Promise<Upload> => {
        try {
          // Add delay between API calls to avoid rate limiting
          await new Promise(resolve => setTimeout(resolve, 100));
          
          const res = await fetch(
            `https://api.linkpreview.net/?key=${API_KEY}&q=${encodeURIComponent(upload.video_link)}`,
            {
              method: 'GET',
              headers: {
                'Content-Type': 'application/json',
              },
            }
          );

          if (!res.ok) {
            throw new Error(`LinkPreview API error: ${res.status}`);
          }

          const meta = await res.json();
          return {
            ...upload,
            previewImage: meta?.image || DEFAULT_IMAGE,
            tags: upload.tags || [],
            category: upload.category || 'General',
          };
        } catch (error) {
          console.warn('Failed to fetch preview for upload:', upload.id, error);
          return {
            ...upload,
            previewImage: DEFAULT_IMAGE,
            tags: upload.tags || [],
            category: upload.category || 'General',
          };
        }
      };

      // Process uploads in smaller batches to avoid overwhelming the API
      const batchSize = 3;
      const withPreviewImages: Upload[] = [];
      
      for (let i = 0; i < uploadsData.length; i += batchSize) {
        const batch = uploadsData.slice(i, i + batchSize);
        const processedBatch = await Promise.all(batch.map(processUpload));
        withPreviewImages.push(...processedBatch);
      }

      // Sort uploads by date (most recent first)
      const sortedUploads = withPreviewImages.sort((a, b) => {
        const dateA = new Date(a.created_at || 0).getTime();
        const dateB = new Date(b.created_at || 0).getTime();
        return dateB - dateA;
      });

      // Calculate user stats
      const now = new Date();
      const weekAgo = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000);
      
      const userStats: UserStats = {
        totalSavedItems: sortedUploads.length,
        recentActivityCount: sortedUploads.filter(upload => {
          const uploadDate = new Date(upload.created_at || 0);
          return uploadDate > weekAgo;
        }).length,
      };

      setState(prev => ({
        ...prev,
        uploads: sortedUploads,
        recentUploads: sortedUploads.slice(0, 10),
        userStats,
        isLoading: false,
        isRefreshing: false,
        hasError: false,
      }));
    } catch (error) {
      console.error('Failed to fetch uploads:', error);
      setState(prev => ({
        ...prev,
        isLoading: false,
        isRefreshing: false,
        hasError: true,
        errorMessage: error instanceof Error ? error.message : 'Failed to load content. Pull down to retry.',
      }));
    }
  }, [userId, API_KEY]);

  // Handle pull-to-refresh
  const handleRefresh = useCallback(() => {
    fetchUploads(true);
  }, [fetchUploads]);

  // Initial data load
  useEffect(() => {
    fetchUploads();
  }, [fetchUploads]);

  return {
    state,
    searchQuery,
    filteredUploads,
    greeting,
    setSearchQuery,
    fetchUploads,
    handleRefresh,
  };
};