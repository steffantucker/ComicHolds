import React, { Component } from "react";
import { connect } from "react-redux";
import { v4 as uuidv4 } from "uuid";

import Listitem from "../Listitem";
import { getComics } from "../../redux"
import AddComic from "../AddComic";

class ComicList extends Component {
	constructor(props) {
		super(props)
		this.props.getComics()
	}
	render() {
		console.log("serving comiclist", this.props.list)
		return (
			<div>
				<div className="displayContainer">
					{this.props.type === "comic" &&
						this.props.list.map(v => (
							<Listitem
								key={uuidv4()}
								click={() => this.props.history.push(`/comics/{v.id}`)}
								name={v.name}
								description={v.description}
								comicid={v.id}
							/>
						))}
				</div>
				<AddComic />
			</div>
		);
	}
}

export default connect(
	state => ({ list: state.results, type: state.type }),
	{ getComics }
)(ComicList)