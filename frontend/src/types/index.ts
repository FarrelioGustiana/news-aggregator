// User Types
export interface User {
  id: string;
  username: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserCredentials {
  username: string;
  password: string;
}

// Auth Types
export interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
}

// Feed Types
export interface Feed {
  id: number;
  name: string;
  url: string;
  lastFetchedAt: string | null;
  createdAt: string;
  updatedAt: string;
  isSubscribed?: boolean; // Frontend specific property
}

// Subscription Types
export interface Subscription {
  id: number;
  userId: string;
  feedId: number;
  createdAt: string;
  updatedAt: string;
  feed?: Feed; // For expanded subscription data
}

// Article Types
export interface Article {
  id: number;
  feedId: number;
  title: string;
  link: string;
  description: string | null;
  pubDate: string | null;
  guid: string | null;
  createdAt: string;
  updatedAt: string;
  feed?: Feed; // For expanded article data
}

// API Response Types
export interface APIResponse<T> {
  data: T;
  message?: string;
  status: number;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  pageSize: number;
  totalPages: number;
}

// Form Types
export interface RegisterFormData {
  username: string;
  password: string;
  confirmPassword: string;
}

export interface LoginFormData {
  username: string;
  password: string;
}

export interface ProfileFormData {
  username?: string;
  currentPassword?: string;
  newPassword?: string;
  confirmNewPassword?: string;
}

export interface FeedFormData {
  name: string;
  url: string;
}
