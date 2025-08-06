# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Lyked** is a mobile application for saving and organizing favorite videos/images from social media platforms in a private media vault. It's built as a monorepo with a React Native + Expo frontend and a Go backend.

## Architecture

### Frontend (mobile/)

- **React Native + Expo**: Mobile app with tab navigation
- **NativeWind**: Tailwind CSS styling for React Native
- **Context API**: User management via `src/app/context/UserContext.tsx`
- **Error Boundaries**: Wrapped around all main screens in `App.tsx`
- **Environment Variables**: Uses `EXPO_PUBLIC_*` prefixed variables in `.env`

### Backend (backend/)

- **Go + Gin**: REST API server with middleware stack
- **Dual Database**: MongoDB for media uploads/folders, PostgreSQL (GORM) for user data
- **MongoDB v2 Driver**: All imports use `go.mongodb.org/mongo-driver/v2/*` paths
- **Modular Structure**: Handlers, routes, models, utilities separated

### Key Data Flow

1. Users authenticate via UserContext (currently hardcoded UUID)
2. Media uploads stored in MongoDB with user_id reference
3. LinkPreview.net API generates thumbnails for saved links
4. All API calls go through error boundary handling

## Development Commands

### Mobile Development

```bash
cd mobile/
npm start           # Start Expo dev server
npm run ios         # iOS simulator
npm run android     # Android emulator
npm run lint        # ESLint + Prettier check
npm run format      # Auto-fix linting issues
```

### Backend Development

```bash
cd backend/
go run cmd/main.go  # Start server on localhost:3000
go build cmd/main.go # Build binary
go mod tidy         # Clean dependencies
```

## Important Implementation Details

### MongoDB ObjectID Handling

- Always use `bson.ObjectIDFromHex()` for string-to-ObjectID conversion
- Delete operations require proper ObjectID conversion, not string matching
- All models use `bson.ObjectID` type, not `primitive.ObjectID`

### CORS Configuration

Server is configured for Expo development with specific allowed origins:

- `http://localhost:8081` (Metro bundler)
- `exp://192.168.*:8081` (Expo dev client)

### Error Handling Patterns

- Backend: All handlers return structured JSON with `error` and `details` fields
- Frontend: ErrorBoundary components catch React errors with fallback UI
- API calls include proper HTTP status code checking and timeout handling

### Environment Variables

- Backend: Uses `utils.GetEnv()` helper with fallbacks
- Frontend: Expo public variables must be prefixed with `EXPO_PUBLIC_`
- LinkPreview API key should be secured in environment, not hardcoded

### Database Models

- **LykedUploads**: MongoDB collection for media items with user_id, title, description, video_link, tags, folders
- **Folders**: MongoDB collection for organization with user_id reference
- **User**: PostgreSQL table (currently unused in favor of hardcoded context)

## Code Quality Standards

- Backend: Go modules with proper error wrapping using `fmt.Errorf`
- Frontend: ESLint + Prettier enforced, React Hook dependency arrays required
- Both: No hardcoded credentials, environment variable usage mandatory

## API Endpoints

- `POST /uploads/upload` - Create new media item
- `GET /uploads/all?user_id=<uuid>` - Fetch user's uploads
- `DELETE /uploads/delete?id=<objectid>` - Remove upload by MongoDB ObjectID
