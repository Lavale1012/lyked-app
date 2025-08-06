import { Upload, UserStats, GreetingType } from '../types/HomeScreen.types';

/**
 * Generate a greeting based on the current time
 */
export const getTimeBasedGreeting = (): string => {
  const hour = new Date().getHours();
  
  if (hour < 6) return 'Working Late? 🌙';
  if (hour < 12) return 'Good Morning! ☀️';
  if (hour < 18) return 'Good Afternoon! 🌤️';
  return 'Good Evening! 🌅';
};

/**
 * Calculate user statistics from uploads
 */
export const calculateUserStats = (uploads: Upload[]): UserStats => {
  const now = new Date();
  const weekAgo = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000);
  
  return {
    totalSavedItems: uploads.length,
    recentActivityCount: uploads.filter(upload => {
      const uploadDate = new Date(upload.created_at || 0);
      return uploadDate > weekAgo;
    }).length,
  };
};

/**
 * Filter uploads based on search query
 */
export const filterUploads = (uploads: Upload[], searchQuery: string): Upload[] => {
  if (!searchQuery.trim()) return uploads;
  
  const query = searchQuery.toLowerCase();
  return uploads.filter(upload =>
    upload.title.toLowerCase().includes(query) ||
    upload.description.toLowerCase().includes(query) ||
    upload.tags?.some(tag => tag.toLowerCase().includes(query)) ||
    upload.category?.toLowerCase().includes(query)
  );
};

/**
 * Sort uploads by date (most recent first)
 */
export const sortUploadsByDate = (uploads: Upload[]): Upload[] => {
  return uploads.sort((a, b) => {
    const dateA = new Date(a.created_at || 0).getTime();
    const dateB = new Date(b.created_at || 0).getTime();
    return dateB - dateA;
  });
};

/**
 * Process upload with fallback preview image
 */
export const processUploadWithFallback = (
  upload: Upload, 
  defaultImage: string
): Upload => ({
  ...upload,
  previewImage: upload.previewImage || defaultImage,
  tags: upload.tags || [],
  category: upload.category || 'General',
});

/**
 * Batch process uploads to avoid overwhelming APIs
 */
export const batchProcess = <T>(
  items: T[], 
  batchSize: number = 3
): T[][] => {
  const batches: T[][] = [];
  for (let i = 0; i < items.length; i += batchSize) {
    batches.push(items.slice(i, i + batchSize));
  }
  return batches;
};

/**
 * Add delay between operations
 */
export const delay = (ms: number): Promise<void> => 
  new Promise(resolve => setTimeout(resolve, ms));

/**
 * Generate card dimensions based on screen width
 */
export const calculateCardDimensions = (
  screenWidth: number, 
  numColumns: number = 2, 
  margin: number = 16
) => ({
  cardWidth: (screenWidth - (numColumns + 1) * margin) / numColumns,
  cardHeight: function(this: { cardWidth: number }) { 
    return this.cardWidth * 0.75; 
  },
  imageHeight: function(this: { cardWidth: number }) { 
    return this.cardWidth * 0.75 * 0.7; 
  },
});

/**
 * Validate environment variables
 */
export const validateEnvironment = () => {
  const apiKey = process.env.EXPO_PUBLIC_LINKPREVIEW_API_KEY;
  const baseUrl = process.env.EXPO_PUBLIC_API_BASE_URL;
  
  return {
    isValid: !!(apiKey && baseUrl),
    apiKey,
    baseUrl: baseUrl || 'http://localhost:8080',
    errors: [
      !apiKey && 'LinkPreview API key is missing',
      !baseUrl && 'API base URL is missing',
    ].filter(Boolean) as string[],
  };
};

export default {
  getTimeBasedGreeting,
  calculateUserStats,
  filterUploads,
  sortUploadsByDate,
  processUploadWithFallback,
  batchProcess,
  delay,
  calculateCardDimensions,
  validateEnvironment,
};