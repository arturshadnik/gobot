"use client";

import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  ReactNode,
} from "react";
import { auth } from "@/config/firebaseConfig";
import { getRedirectResult, onAuthStateChanged, User } from "firebase/auth";

interface AuthContextProps {
  user: User | null;
  userLoading: boolean;
  signOut: () => Promise<void>;
}

interface AuthProviderProps {
  children: ReactNode;
}

const AuthContext = createContext<AuthContextProps>({
  user: null,
  userLoading: true,
  signOut: async () => {},
});

const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [userLoading, setUserLoading] = useState<boolean>(true);

  useEffect(() => {
    getRedirectResult(auth)
      .then(() => {
        const unsubscribe = onAuthStateChanged(auth, (user) => {
          setUser(user);
          setUserLoading(false);
        });
        return () => {
          unsubscribe();
        };
      })
      .catch((error: any) => {
        console.error(error);
        setUserLoading(false);
        setUser(null);
      });
  }, []);

  const signOut = async () => {
    try {
      await auth.signOut();
      console.log("Signed out")
      setUser(null);
    } catch (error: any) {
      console.error("Error signing out: ", error);
    }
  };

  return (
    <AuthContext.Provider value={{ user, userLoading, signOut }}>
      {children}
    </AuthContext.Provider>
  );
};

const useAuth = () => useContext(AuthContext);

export { AuthProvider, useAuth };
