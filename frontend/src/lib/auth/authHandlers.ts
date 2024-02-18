import { signInWithRedirect, GoogleAuthProvider, getRedirectResult } from "firebase/auth";
import { auth } from "@/config/firebaseConfig";
import nookies from "nookies"

const handleGoogleLogin = async () => {
    const provider = new GoogleAuthProvider();

    try {
        await signInWithRedirect(auth, provider)
        const results = await getRedirectResult(auth);

        const token = await results?.user.getIdToken();
        const displayName = results?.user.displayName || "";

        if (token) {
            nookies.set(undefined, "token", token, {path: "/"});
            nookies.set(undefined, "name", displayName, {path: "/"});
        } else {
            console.error("No token found")
        }
        return true;
    } catch (error: any) {
        console.error(error)
        return false;
    }
};

export default handleGoogleLogin;

