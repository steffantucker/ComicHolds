import { createStore, applyMiddleware } from "redux";
import axios from "axios";
import thunk from "redux-thunk"

const initState = {
	type: null,
	result: []
};

export const getComics = () => {
	return function (dispatch) {
		axios
			.get(`/comics`)
			.then(comics => {
				console.log("comics result", comics.data)
				dispatch({
					type: "SHOW_COMICS",
					data: comics.data
				});
			})
			.catch(err => console.error(err));
	};
};

export const addComic = (data) => {
	return function (dispatch) {
		axios
			.post(`/comics`, data)
			.then(id => {
				data.id = id.data
				dispatch({
					type: "ADD_COMIC",
					data: data
				});
			})
			.catch(err => console.error(err))
	};
};

const reducer = (prev = initState, action) => {
	let { type, results } = prev;
	switch (action.type) {
		case "SHOW_COMICS":
			type = "comic";
			results = action.data;
			break;
		case "ADD_COMIC":
			type = "newcomic";
			results = action.data;
			break;
		default:
			type = null;
	}
	return { type, results };
}

export default createStore(
	reducer,
	window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__(),
	applyMiddleware(thunk)
)