export interface Upload {
  id: string;
  title: string;
  description: string;
  video_link: string;
  user_id: string;
  previewImage: string;
  created_at?: string;
  updated_at?: string;
  category?: string;
  tags?: string[];
}

export interface UserStats {
  totalSavedItems: number;
  recentActivityCount: number;
  favoritesCount?: number;
}

export interface HomeScreenState {
  uploads: Upload[];
  recentUploads: Upload[];
  isLoading: boolean;
  isRefreshing: boolean;
  hasError: boolean;
  errorMessage: string;
  userStats: UserStats;
  debugInfo?: any;
}

export type GreetingType = 'morning' | 'afternoon' | 'evening' | 'night';

export interface HeaderProps {
  greeting: string;
  userStats: UserStats;
  onSearch: (query: string) => void;
  onFilterPress: () => void;
}

export interface MediaCardProps {
  upload: Upload;
  width: number;
  onPress: (upload: Upload) => void;
  onShare: (upload: Upload) => void;
  onOrganize: (upload: Upload) => void;
}

export interface RecentlyAddedProps {
  uploads: Upload[];
  onSeeAll: () => void;
  onItemPress: (upload: Upload) => void;
}

export interface LoadingSkeletonProps {
  width: number;
  height: number;
  borderRadius?: number;
}