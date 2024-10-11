import { combineReducers } from "redux";
import {
  SHORTEN_URL_REQUEST,
  SHORTEN_URL_SUCCESS,
  SHORTEN_URL_FAILURE,
  GET_ORIGINAL_URL_REQUEST,
  GET_ORIGINAL_URL_SUCCESS,
  GET_ORIGINAL_URL_FAILURE,
} from "./actions";

const initialState = {
  shortUrl: "",
  originalUrl: "",
  loading: false,
  error: null,
};

const urlReducer = (state = initialState, action) => {
  switch (action.type) {
    case SHORTEN_URL_REQUEST:
    case GET_ORIGINAL_URL_REQUEST:
      return { ...state, loading: true, error: null };
    case SHORTEN_URL_SUCCESS:
      return { ...state, shortUrl: action.payload, loading: false };
    case GET_ORIGINAL_URL_SUCCESS:
      return { ...state, originalUrl: action.payload, loading: false };
    case SHORTEN_URL_FAILURE:
    case GET_ORIGINAL_URL_FAILURE:
      return { ...state, error: action.payload, loading: false };
    default:
      return state;
  }
};

export default combineReducers({
  url: urlReducer,
});
