import axios from "axios"

export const Navlinks = [
    {id: 1, name: "Inbox", link:"/", getCount:"inbox_num"},
    {id: 2, name: "Sent Box", link:"/sentbox", getCount: "sentbox_num"},
    {id: 3, name: "Favourite", link:"/favourite", getCount: "fav_num"},
]

export const customFetch = axios.create({
    baseURL: "http://localhost:8080/api/v1"
})

function debounceOuter(){
    let timer

    return (cb) => {
        clearTimeout(timer)
        timer = setTimeout(cb, 1000)
    } 
}

export const debounce = debounceOuter()