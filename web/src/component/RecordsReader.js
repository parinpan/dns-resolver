import React from 'react'
import RecordFetcher from "../api/RecordFetcher"

import {
    Table,
    TableContainer,
    TableRow,
    TableCell,
    TableBody, Typography
} from "@mui/material"

export default class RecordsReader extends React.Component {
    constructor(props) {
        super(props);
        this.state = {loading: false, data: {}}
        this.fetcher = new RecordFetcher('http://localhost')
    }

    componentDidUpdate(prevProps) {
        const {queryHostname, queryRecord} = this.props
        const self = this

        if (prevProps.queryHostname !== queryHostname || prevProps.queryRecord !== queryRecord) {
            this.fetcher.fetch(queryHostname, queryRecord)
                .then(response => {
                    self.setState({data: response.data})
                })
                .catch(e => {
                    console.log(e)
                })
        }
    }

    renderItems(items) {
        return items.map((item, i) => {
            return (
                <div key={i + item.key}>
                    <Typography variant="subtitle2">
                        <strong>{item.key}</strong>
                    </Typography>
                    <Typography variant="subtitle2">
                        {item.value}
                    </Typography>
                </div>
            )
        })
    }

    renderRecords(recordTypes, records) {
        return recordTypes.map(key => {
            return (
                <TableRow key={key}>
                    <TableCell>{key}</TableCell>
                    <TableCell>
                        {records[key].map((items, i) => {
                            return (
                                <div key={i} className={"sub-items"}>
                                    {this.renderItems(items)}
                                </div>
                            )
                        })}
                    </TableCell>
                </TableRow>
            )
        })
    }

    render() {
        const records = this.state.data || {}
        const recordTypes = Object.keys(records)

        if (this.props.queryHostname === "" || this.props.queryRecord === "") {
            return (
                <Typography className={"records-plain-text"} variant="subtitle2">
                    <strong>Please enter a hostname and select the record type</strong>
                </Typography>
            )
        }

        if (recordTypes.length === 0) {
            return (
                <Typography className={"records-plain-text"} variant="subtitle2">
                    <strong>Records not found!</strong>
                </Typography>
            )
        }

        return (
            <TableContainer>
                <Table>
                    <TableBody>
                        {this.renderRecords(recordTypes, records)}
                    </TableBody>
                </Table>
            </TableContainer>
        )
    }
}
