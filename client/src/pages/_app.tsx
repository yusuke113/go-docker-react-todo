import '@/styles/globals.css'
import { CsrfToken } from '@/types'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import axios from 'axios'
import type { AppProps } from 'next/app'
import { useEffect } from 'react'

const queryClient = new QueryClient({})

export default function App({ Component, pageProps }: AppProps) {
  useEffect(() => {
    const getCsrfToken = async () => {
      axios.defaults.withCredentials = true
      const { data } = await axios.get<CsrfToken>(
        `${process.env.NEXT_PUBLIC_APP_API_URL}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
      console.log(data)
    }
    getCsrfToken()
  }, [])

  return (
    <QueryClientProvider client={queryClient}>
      <Component {...pageProps} />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  )
}
