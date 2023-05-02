import { CsrfToken } from "@/types"
import axios from "axios"
import { useRouter } from "next/router"
import { useStore } from "zustand"


export const useError = () => {
  const resetEditedTask = useStore((state) => state.resetEditedTask)
  const getCsrfToken = async () => {
    const { data } = await axios.get<CsrfToken>(
      `${process.env.NEXT_PUBLIC_APP_API_URL}/csrf`
    )
    axios.defaults.headers.common["X-CSRF-TOKEN"] = data.csrf_token
  }
  const switchErrorHandling = (msg: string) => {
    const router = useRouter()

    switch (msg) {
      case 'invalid csrf token':
        getCsrfToken()
        alert('CSRF token is invalid, please try again')
        break
      case 'invalid or expired jwt':
        alert('access token expired, please login')
        resetEditedTask()
        router.replace('/') // ここでリダイレクト
        break
      case 'missing or malformed jwt':
        alert('access token is not valid, please login')
        resetEditedTask()
        router.replace('/') // ここでリダイレクト
        break
      case 'duplicated key not allowed':
        alert('email already exist, please use another one')
        break
      case 'crypto/bcrypt: hashedPassword is not the hash of the given password':
        alert('password is not correct')
        break
      case 'record not found':
        alert('email is not correct')
        break
      default:
        alert(msg)
    }
  }
  return { switchErrorHandling }
}