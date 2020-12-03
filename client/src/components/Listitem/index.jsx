import React from "react";

const Listitem = props => {
	return (
		<div className="comic" id={props.id} onClick={props.click ? props.click : null}>
			<h5>{props.name}</h5>
			{props.description ? (
				<p className="comic-description">{props.description}</p>
			) : null}
		</div>
	)
}

export default Listitem;