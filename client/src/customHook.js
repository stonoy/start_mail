import axios from "axios"
import { useEffect, useState } from "react"
import { customFetch } from "./utils"
import { toast } from "react-toastify"

export const useCustomHook = (url, token) => {
    const [data, setData] = useState(null)

    useEffect(() => {
        const fetchData = async () => {
            try {
                const resp = await customFetch.get(`/${url}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })
    
                setData(resp?.data)
            } catch (error) {
               toast.error(error?.response?.data?.Msg) 
            }
        }

        fetchData()

    }, [])

    return {data}
}