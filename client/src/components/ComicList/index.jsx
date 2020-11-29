import React, { Component } from "react";
import { connect } from "react-redux";
import uuidv4 from "uuid/v4";

import Listitem from "../Listitem";
import { getComics } from "../../redux"

class ComicList extends Component {
	render() {
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
							/>
						))}
				</div>
			</div >
		);
	}
}

export default connect(
	state => ({ list: state.results, type: state.type }),
	{ getComics }
)(ComicList)