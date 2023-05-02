import { useRouter } from "next/router"
import { useStore } from "zustand"
import { useError } from "./useError"
import axios from "axios"
import { useMutation } from "@tanstack/react-query"

export const useMutateAuth = () => {
  const router = useRouter()
  const resetEditedTask = useStore((state) => state.resetEditedTask)
  const { switchErrorHandling } = useError()
  const loginMutation = useMutation(
    async (user: Credential) =>
      await axios.post(`${process.env.NEXT_PUBLIC_APP_API_URL}/login`, user),
    {
      onSuccess: () => {
        router.push('/todos')
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )

  const registerMutation = useMutation(
    async (user: Credential) =>
      await axios.post(`${process.env.NEXT_PUBLIC_APP_API_URL}/signup`, user),
    {
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )

  const logoutMutation = useMutation(
    async () => await axios.post(`${process.env.NEXT_PUBLIC_APP_API_URL}/logout`),
    {
      onSuccess: () => {
        resetEditedTask()
        router.push('/')
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message)
        } else {
          switchErrorHandling(err.response.data)
        }
      },
    }
  )

  return { loginMutation, registerMutation, logoutMutation }
}