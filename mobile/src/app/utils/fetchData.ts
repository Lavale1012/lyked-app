export interface ApiResponse<T = any> {
  data?: T;
  error?: {
    status: number;
    statusText: string;
    message: string;
    url: string;
    body?: string;
    headers?: Record<string, string>;
  };
}

export interface FetchOptions {
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE';
  headers?: Record<string, string>;
  body?: string;
  enableDetailedLogging?: boolean;
}

export const fetchWithLogging = async <T = any>(
  url: string, 
  options: FetchOptions = {}
): Promise<ApiResponse<T>> => {
  const { 
    method = 'GET', 
    headers = {}, 
    body, 
    enableDetailedLogging = true 
  } = options;

  // Add Origin header for local API calls to match backend CORS requirements
  const baseHeaders: Record<string, string> = {
    'Content-Type': 'application/json',
  };
  
  // If this is a call to localhost, add the Origin header
  if (url.includes('localhost') || url.includes('127.0.0.1')) {
    const apiBaseUrl = process.env.EXPO_PUBLIC_API_BASE_URL || 'http://localhost:8080';
    baseHeaders['Origin'] = apiBaseUrl;
  }

  const requestConfig: RequestInit = {
    method,
    headers: {
      ...baseHeaders,
      ...headers,
    },
  };

  if (body) {
    requestConfig.body = body;
  }

  if (enableDetailedLogging) {
    console.group(`🌐 API Request: ${method} ${url}`);
    console.log('📍 Full URL:', url);
    console.log('🔧 Config:', requestConfig);
    console.log('📋 Headers:', requestConfig.headers);
    if (body) console.log('📦 Body:', body);
    console.groupEnd();
  }

  try {
    const startTime = Date.now();
    const response = await fetch(url, requestConfig);
    const duration = Date.now() - startTime;

    // Convert headers to object for logging
    const responseHeaders: Record<string, string> = {};
    response.headers.forEach((value, key) => {
      responseHeaders[key] = value;
    });

    if (enableDetailedLogging) {
      console.group(`📡 API Response: ${response.status} ${response.statusText} (${duration}ms)`);
      console.log('📍 URL:', url);
      console.log('📊 Status:', response.status, response.statusText);
      console.log('📋 Headers:', responseHeaders);
      console.log('⏱️ Duration:', `${duration}ms`);
    }

    if (!response.ok) {
      let errorBody = '';
      try {
        errorBody = await response.text();
        if (enableDetailedLogging) {
          console.log('❌ Error Body:', errorBody);
        }
      } catch (bodyError) {
        if (enableDetailedLogging) {
          console.log('❌ Could not read error body:', bodyError);
        }
      }

      if (enableDetailedLogging) {
        console.groupEnd();
      }

      return {
        error: {
          status: response.status,
          statusText: response.statusText,
          message: `HTTP ${response.status}: ${response.statusText}`,
          url,
          body: errorBody,
          headers: responseHeaders,
        },
      };
    }

    const data = await response.json();
    
    if (enableDetailedLogging) {
      console.log('✅ Response Data:', data);
      console.groupEnd();
    }

    return { data };
  } catch (error) {
    if (enableDetailedLogging) {
      console.group('💥 Network/Parse Error');
      console.log('📍 URL:', url);
      console.error('❌ Error:', error);
      console.log('🔍 Error details:', {
        name: error instanceof Error ? error.name : 'Unknown',
        message: error instanceof Error ? error.message : String(error),
        stack: error instanceof Error ? error.stack : undefined,
      });
      console.groupEnd();
    }

    return {
      error: {
        status: 0,
        statusText: 'Network Error',
        message: error instanceof Error ? error.message : 'Network request failed',
        url,
      },
    };
  }
};

// Backward compatibility
export const fetch_Data = async (url: string) => {
  const result = await fetchWithLogging(url);
  if (result.error) {
    throw new Error(`${result.error.message} - Status: ${result.error.status}`);
  }
  return result.data;
};
