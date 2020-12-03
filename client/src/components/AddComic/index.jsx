import React, { Component } from "react";
import { connect } from "react-redux";

import { addComic } from "../../redux";

class AddComic extends Component {
	render() {
		return (
			<div>
				
			</div>
		);
	}
}

export default connect(
	null,
	{ addComic }
)(AddComic)