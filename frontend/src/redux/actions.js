export const SHORTEN_URL_REQUEST = "SHORTEN_URL_REQUEST";
export const SHORTEN_URL_SUCCESS = "SHORTEN_URL_SUCCESS";
export const SHORTEN_URL_FAILURE = "SHORTEN_URL_FAILURE";
export const GET_ORIGINAL_URL_REQUEST = "GET_ORIGINAL_URL_REQUEST";
export const GET_ORIGINAL_URL_SUCCESS = "GET_ORIGINAL_URL_SUCCESS";
export const GET_ORIGINAL_URL_FAILURE = "GET_ORIGINAL_URL_FAILURE";

export const shortenUrl = (url) => ({
  type: SHORTEN_URL_REQUEST,
  payload: url,
});
export const getOriginalUrl = (code, checkOnly = false) => ({
  type: GET_ORIGINAL_URL_REQUEST,
  payload: { code, checkOnly },
});
