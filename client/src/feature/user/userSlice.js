import {createSlice, createAsyncThunk} from "@reduxjs/toolkit"
import { customFetch } from "../../utils"
import { toast } from "react-toastify"

const initialState = {
    submitting: false,
    user: null,
    token: ""
}

export const loginUser = createAsyncThunk("user/login", 
    async (data, thunkAPI) => {
       
        try {
            const resp = await customFetch.post("/login", data)
            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response?.data?.Msg)
        }
    }
)

export const registerUser = createAsyncThunk("user/register", 
    async (data, thunkAPI) => {
       
        try {
            const resp = await customFetch.post("/register", data)
            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response?.data?.Msg)
        }
    }
)

const userSlice = createSlice({
    name: "user",
    initialState: JSON.parse(localStorage.getItem("user")) || initialState,
    reducers: {
        logout : (state, {payload}) => {
            return initialState
        },
    },
    extraReducers: (builder) => {
        builder.addCase(loginUser.pending, (state, {payload}) => {
            state.submitting = true
        }).addCase(loginUser.fulfilled, (state, {payload}) => {
            state.submitting = false
            state.token = payload.token
            state.user = payload.user
            localStorage.setItem("user", JSON.stringify(state))
            toast.success("logged In!")
        }).addCase(loginUser.rejected, (state, {payload}) => {
            state.submitting = false
            toast.error(payload)
        }).addCase(registerUser.pending, (state, {payload}) => {
            state.submitting = true
        }).addCase(registerUser.fulfilled, (state, {payload}) => {
            state.submitting = false
        }).addCase(registerUser.rejected, (state, {payload}) => {
            state.submitting = false
            toast.error(payload)
        })
    }
})

export const {logout} = userSlice.actions

export default userSlice.reducer