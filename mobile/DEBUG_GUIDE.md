# Mobile App Debugging Guide

This guide explains the enhanced error handling and debugging features added to help diagnose the 500 error and other API issues.

## Features Added

### 1. Enhanced Error Logging (`fetchWithLogging` utility)
- **Location**: `/src/app/utils/fetchData.ts`
- **Purpose**: Provides comprehensive logging of all API requests and responses
- **Features**:
  - Logs full request URL, headers, and body
  - Logs response status, headers, and body (including error bodies)
  - Measures and logs request duration
  - Groups console logs for better readability
  - Captures and formats error details

### 2. Improved Error UI (`ErrorDebugView` component)
- **Location**: `/src/app/components/ErrorDebugView.tsx`
- **Purpose**: Shows detailed error information to developers
- **Features**:
  - User-friendly error messages with context
  - Expandable debug details
  - Share debug info functionality
  - Quick info panel showing key details (URL, status, user ID, etc.)
  - Different icons for different error types

### 3. Environment Validation
- **Location**: Enhanced in `HomeScreen.tsx`
- **Purpose**: Validates and logs environment configuration
- **Features**:
  - Checks if API_BASE_URL is set
  - Checks if LinkPreview API key is set
  - Logs configuration details on app startup
  - Shows helpful error messages for missing config

### 4. Debug Panel (Development Only)
- **Location**: `/src/app/components/DebugPanel.tsx`
- **Purpose**: Allows testing different error scenarios
- **Features**:
  - Only visible in development mode (`__DEV__`)
  - Red bug button in bottom left corner
  - Test scenarios for different error types
  - Console logging for verification

## How to Debug the 500 Error

### Step 1: Check Console Logs
1. Open your development console (Metro bundler terminal or React Native debugger)
2. Refresh the HomeScreen or pull to refresh
3. Look for grouped console logs starting with:
   - 🔧 **Environment Configuration** - Shows your API settings
   - 🌐 **API Request** - Shows the exact request being made
   - 📡 **API Response** - Shows the server response details

### Step 2: Check Environment Variables
The app will automatically validate and log your environment setup:
```
🔧 Environment Configuration
📍 API_BASE_URL: http://localhost:8080
🔑 API_KEY: Set (ad738c4c...)
👤 Current User ID will be logged when available
```

Ensure:
- `EXPO_PUBLIC_API_BASE_URL` points to your backend server
- `EXPO_PUBLIC_LINKPREVIEW_API_KEY` is set

### Step 3: Analyze the Error Details
If you get a 500 error, the enhanced UI will show:
- **Quick Info Panel**: URL being called, HTTP status, User ID, API base URL
- **Error Details** (expandable): Full request/response details, headers, error body
- **Share Function**: Export all debug info for sharing with backend developers

### Step 4: Test with Debug Panel
1. Look for the red 🐛 button in the bottom left (development only)
2. Tap it to open the debug panel
3. Test different error scenarios to verify error handling works correctly

## Key Information for Backend Debugging

When you encounter the 500 error, the logs will show:

### Request Details
```
🌐 API Request: GET http://localhost:8080/uploads/all?user_id=7a5e1f2e-8d33-4abc-b8cb-9f9fdab6f1df
📍 Full URL: http://localhost:8080/uploads/all?user_id=7a5e1f2e-8d33-4abc-b8cb-9f9fdab6f1df
🔧 Config: {method: 'GET', headers: {...}}
📋 Headers: {Content-Type: 'application/json'}
```

### Response Details
```
📡 API Response: 500 Internal Server Error (1234ms)
📍 URL: http://localhost:8080/uploads/all?user_id=7a5e1f2e-8d33-4abc-b8cb-9f9fdab6f1df
📊 Status: 500 Internal Server Error
📋 Headers: {content-type: 'application/json', ...}
❌ Error Body: {"error": "Database connection failed", "stack": "..."}
⏱️ Duration: 1234ms
```

## Configuration Files

### Current Environment (.env)
```
EXPO_PUBLIC_LINKPREVIEW_API_KEY=ad738c4cf4e7548994a8e95ef192859c
EXPO_PUBLIC_API_BASE_URL=http://localhost:8080
```

### User Context
- Current hardcoded user ID: `7a5e1f2e-8d33-4abc-b8cb-9f9fdab6f1df`
- Location: `/src/app/context/UserContext.tsx`

## Tips for Development

1. **Keep Console Open**: The detailed logs are crucial for debugging
2. **Use Pull to Refresh**: Easy way to retry API calls and see fresh logs
3. **Check Network**: Ensure your backend server is running on the configured port
4. **Share Debug Info**: Use the built-in share function to export error details
5. **Test Error Scenarios**: Use the debug panel to verify error handling works

## Files Modified

- ✅ `src/app/utils/fetchData.ts` - Enhanced API utility with logging
- ✅ `src/app/screens/HomeScreen.tsx` - Improved error handling and logging
- ✅ `src/app/types/HomeScreen.types.ts` - Added debugInfo field
- ✅ `src/app/components/ErrorDebugView.tsx` - New detailed error UI
- ✅ `src/app/components/DebugPanel.tsx` - New development debug panel

## Next Steps

1. Start your backend server
2. Ensure it's running on the correct port (8080 or update the env variable)
3. Test the API endpoint directly: `GET http://localhost:8080/uploads/all?user_id=7a5e1f2e-8d33-4abc-b8cb-9f9fdab6f1df`
4. Check the mobile app logs for the exact error details from the server
5. Share the error body and stack trace with the backend team

The enhanced error handling will now provide all the information needed to diagnose whether the issue is:
- Configuration (wrong API URL)
- Network connectivity
- Backend server error (500 response with details)
- Data parsing issues