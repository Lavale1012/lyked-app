import React, { createContext, useContext, ReactNode } from 'react';

type UserContextType = {
  userId: string;
};

const UserContext = createContext<UserContextType | undefined>(undefined);

type UserProviderProps = {
  children: ReactNode;
};

export const UserProvider: React.FC<UserProviderProps> = ({ children }) => {
  // TODO: Replace with actual authentication logic
  const userId = '7a5e1f2e-8d33-4abc-b8cb-9f9fdab6f1df';

  return <UserContext.Provider value={{ userId }}>{children}</UserContext.Provider>;
};

export const useUser = (): UserContextType => {
  const context = useContext(UserContext);
  if (context === undefined) {
    throw new Error('useUser must be used within a UserProvider');
  }
  return context;
};
