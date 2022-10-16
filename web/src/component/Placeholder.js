import {Typography} from "@mui/material";
import React from "react";

export default class Placeholder extends React.Component {
    render() {
        return (
            <Typography className={"records-plain-text"} variant="subtitle2">
                <strong>{this.props.text}</strong>
            </Typography>
        );
    }
}
