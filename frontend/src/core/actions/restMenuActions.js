export const sendMenu = (data) => {
    return {
        type: 'SEND_MENU',
        payload: data 
    } 
}
export const SEND_MENU = 'SEND_MENU'

export const getMenu = (data) => {
    return {
        type: 'GET_MENU',
        payload: data 
    } 
}
export const GET_MENU = 'GET_MENU'

export const getMenuSuccess = (data) => {
    return {
        type: 'GET_MENU_SUCCESS',
        payload: data 
    } 
}
export const getMenuError = (data) => {
    return {
        type: 'GET_MENU_ERROR',
        payload: data 
    } 
}