import React from "react"
import {Box, Button} from '@mui/material'

const boxStyle = {
    margin: "20px 0px",
    display: 'inline-flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
    alignItems: 'center',
}

export default class RecordsMenu extends React.Component {
    constructor(props) {
        super(props)
        this.records = props.records
        this.buttonRef = React.createRef()
        this.state = {active: "ANY"}
    }

    isActive(record) {
        return this.state.active === record
    }

    render() {
        const {changeRecord} = this.props

        return (
            <Box sx={boxStyle}>
                {this.props.records.map(record =>
                    <Button
                        className={"records-menu"}
                        ref={this.buttonRef}
                        key={record}
                        onClick={(_) => {
                            this.setState({active: record})
                            changeRecord(record)
                        }}
                        variant={this.isActive(record) ? "contained" : "outlined"}>{record}</Button>
                )}
            </Box>
        )
    }
}
