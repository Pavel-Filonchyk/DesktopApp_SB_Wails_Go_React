const initialState = {
    menu: [],
    postDishes: [],
    token: ''
}

const restMenuReducer = (state = initialState, action) => {
    switch (action.type){ 
        case 'SEND_MENU':
           
            return {
                ...state,
                menu: action.payload
            }
        case 'POST_DISHES':
            console.log(action.payload)
            return {
                ...state,
                postDishes: state.menu[action.payload]
            }
        default: 
        return state;  
    }
}

export default restMenuReducer

