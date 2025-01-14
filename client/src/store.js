import {configureStore} from "@reduxjs/toolkit"
import userReducer from "./feature/user/userSlice"
import emailReducer from "./feature/email/emailSlice"

export const store = configureStore({
    reducer : {
        user: userReducer,
        email: emailReducer,
    }
})