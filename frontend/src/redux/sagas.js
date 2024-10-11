import { call, put, takeLatest } from "redux-saga/effects";
import {
  SHORTEN_URL_REQUEST,
  SHORTEN_URL_SUCCESS,
  SHORTEN_URL_FAILURE,
  GET_ORIGINAL_URL_REQUEST,
  GET_ORIGINAL_URL_SUCCESS,
  GET_ORIGINAL_URL_FAILURE,
} from "./actions";

function* shortenUrlSaga(action) {
  try {
    const response = yield call(fetch, "http://localhost:8080/post-url", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ url: action.payload }),
    });
    const data = yield response.json();
    if (data.code === 200) {
      yield put({ type: SHORTEN_URL_SUCCESS, payload: data.model.url });
    } else {
      yield put({ type: SHORTEN_URL_FAILURE, payload: data.msg });
    }
  } catch (error) {
    yield put({ type: SHORTEN_URL_FAILURE, payload: error.message });
  }
}

function* getOriginalUrlSaga(action) {
  try {
    const { code, checkOnly } = action.payload;
    const url = `http://localhost:8080/get-url?code=${code}${
      checkOnly ? "_" : ""
    }`;
    const response = yield call(fetch, url);
    const data = yield response.json();
    if (data.code === 200 || data.code === 201) {
      yield put({ type: GET_ORIGINAL_URL_SUCCESS, payload: data.model.url });
    } else {
      yield put({ type: GET_ORIGINAL_URL_FAILURE, payload: data.msg });
    }
  } catch (error) {
    yield put({ type: GET_ORIGINAL_URL_FAILURE, payload: error.message });
  }
}

export default function* rootSaga() {
  yield takeLatest(SHORTEN_URL_REQUEST, shortenUrlSaga);
  yield takeLatest(GET_ORIGINAL_URL_REQUEST, getOriginalUrlSaga);
}
