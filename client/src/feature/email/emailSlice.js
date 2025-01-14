import { createAsyncThunk, createSlice } from "@reduxjs/toolkit"
import { toast } from "react-toastify"
import { customFetch } from "../../utils"


const initialState = {
    emailCounts: null,
    countsLoading: false,
    data: null,
    theEmail: null,
    loading: false,
    submitting: false,
    fav: null,
    page: 1,
    numOfPages: 1,
}

export const getEmailDetail = createAsyncThunk("email/getEmailDetail", 
    async (emailId, thunkAPI) => {
        
        const {token} = thunkAPI.getState().user
        try {
            const resp = await customFetch.get(`/getemail/${emailId}`, {
                                headers: {
                                    Authorization: `Bearer ${token}`
                                }
                            }) 
                    return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response?.data?.Msg)
        }
    }
)

export const deleteFav = createAsyncThunk("email/deleteFav", 
    async (favId, thunkAPI) => {
        
        const {token} = thunkAPI.getState().user
        try {
            const resp = await customFetch.delete(`/deletefav/${favId}`, {
                                headers: {
                                    Authorization: `Bearer ${token}`
                                }
                            }) 
                    return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response?.data?.Msg)
        }
    }
)

export const writeEmail = createAsyncThunk("email/writeEmail", 
    async (data, thunkAPI) => {
        
        const {token} = thunkAPI.getState().user
        try {
            const resp = await customFetch.post(`/createemails`, data,  {
                                headers: {
                                    Authorization: `Bearer ${token}`
                                }
                            }) 
                    return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response?.data?.Msg)
        }
    }
)

export const getEmailCounts = createAsyncThunk("email/getEmailCounts", 
    async (url, thunkAPI) => {
        
        const {token} = thunkAPI.getState().user
        try {
            const resp = await customFetch.get(`/${url}`, {
                                headers: {
                                    Authorization: `Bearer ${token}`
                                }
                            }) 
                    return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response?.data?.Msg)
        }
    }
)

export const getEmail = createAsyncThunk("email/getEmail", 
    async (url, thunkAPI) => {
        
        const {token} = thunkAPI.getState().user
        try {
            const resp = await customFetch.get(`/${url}`, {
                                headers: {
                                    Authorization: `Bearer ${token}`
                                }
                            }) 
                    return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response?.data?.Msg)
        }
    }
)

export const getFav = createAsyncThunk("email/getFav", 
    async (url, thunkAPI) => {
        
        const {token} = thunkAPI.getState().user
        try {
            const resp = await customFetch.get(`/${url}`, {
                                headers: {
                                    Authorization: `Bearer ${token}`
                                }
                            }) 
                    return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response?.data?.Msg)
        }
    }
)

export const createFav = createAsyncThunk("email/createFav", 
    async (emailId, thunkAPI) => {
        console.log(emailId)
        const {token} = thunkAPI.getState().user
        try {
            const resp = await customFetch.get(`/createFav/${emailId}`, {
                                headers: {
                                    Authorization: `Bearer ${token}`
                                }
                            }) 
                    return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response?.data?.Msg)
        }
    }
)

const emailSlice = createSlice({
    name :"email",
    initialState: JSON.parse(localStorage.getItem("email")) || initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(getEmail.pending, (state, {payload}) => {
            state.loading = true
        }).addCase(getEmail.fulfilled, (state, {payload}) => {
            state.loading = false
            state.data = payload.emails
            state.page = payload.page
            state.numOfPages = payload.numOfPages
            localStorage.setItem("email", JSON.stringify(state))
        }).addCase(getEmail.rejected, (state, {payload}) => {
            state.loading = false
            toast.error(payload)
        }).addCase(getFav.pending, (state, {payload}) => {
            state.loading = true
        }).addCase(getFav.fulfilled, (state, {payload}) => {
            state.loading = false
            state.fav = payload.favourite
            state.page = payload.page
            state.numOfPages = payload.numOfPages
            localStorage.setItem("email", JSON.stringify(state))
        }).addCase(getFav.rejected, (state, {payload}) => {
            state.loading = false
            toast.error(payload)
        }).addCase(createFav.pending, (state, {payload}) => {
            state.submitting = true
        }).addCase(createFav.fulfilled, (state, {payload}) => {
            state.submitting = false
            toast.success("email added to favourite list")
        }).addCase(createFav.rejected, (state, {payload}) => {
            state.submitting = false
            console.log(payload)
            toast.error("already in favoutite email")
        }).addCase(getEmailCounts.pending, (state, {payload}) => {
            state.countsLoading = true
        }).addCase(getEmailCounts.fulfilled, (state, {payload}) => {
            state.countsLoading = false
            state.emailCounts = payload
            localStorage.setItem("email", JSON.stringify(state))
        }).addCase(getEmailCounts.rejected, (state, {payload}) => {
            state.countsLoading = false
            toast.error(payload)
        }).addCase(writeEmail.pending, (state, {payload}) => {
            state.submitting = true
        }).addCase(writeEmail.fulfilled, (state, {payload}) => {
            state.submitting = false
            toast.success("email sent")
        }).addCase(writeEmail.rejected, (state, {payload}) => {
            state.submitting = false
            // console.log(payload)
            toast.error(payload)
        }).addCase(deleteFav.pending, (state, {payload}) => {
            state.submitting = true
        }).addCase(deleteFav.fulfilled, (state, {payload}) => {
            state.submitting = false
            toast.success("email removed from favourite")
        }).addCase(deleteFav.rejected, (state, {payload}) => {
            state.submitting = false
            // console.log(payload)
            toast.error(payload)
        }).addCase(getEmailDetail.pending, (state, {payload}) => {
            state.loading = true
        }).addCase(getEmailDetail.fulfilled, (state, {payload}) => {
            state.loading = false
            state.theEmail = payload?.email
            localStorage.setItem("email", JSON.stringify(state))
        }).addCase(getEmailDetail.rejected, (state, {payload}) => {
            state.loading = false
            toast.error(payload)
        })
    }
})

export default emailSlice.reducer