import { signInWithRedirect, GoogleAuthProvider } from "firebase/auth";
import { auth } from "@/config/firebaseConfig";

const handleGoogleLogin = async () => {
    const provider = new GoogleAuthProvider();

    try {
        await signInWithRedirect(auth, provider)
        return true;
    } catch (error: any) {
        console.error(error)
        return false;
    }
};

export default handleGoogleLogin;

