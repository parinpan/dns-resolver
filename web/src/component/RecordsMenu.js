import React from 'react'
import {Box, Button} from '@mui/material'

const boxStyle = {
    margin: "20px 0px",
    display: 'inline-flex',
    flexWrap: 'wrap',
    justifyContent: 'center'
}

export default class RecordsMenu extends React.Component {
    constructor(props) {
        super(props);
        this.records = props.records
    }

    render() {
        return (
            <Box sx={boxStyle}>
                {this.props.records.map(record =>
                    <Button variant="text">{record}</Button>
                )}
            </Box>
        )
    }
}
