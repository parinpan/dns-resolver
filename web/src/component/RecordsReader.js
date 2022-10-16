import React from "react"
import Placeholder from "./Placeholder"

import {
    Table,
    TableContainer,
    TableRow,
    TableCell,
    TableBody, Typography
} from "@mui/material"


export default class RecordsReader extends React.Component {
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
        const records = this.props.records || {}
        const recordTypes = Object.keys(records)

        if (recordTypes.length === 0) {
            return (<Placeholder text={"Records not found"}/>)
        }

        return (
            <TableContainer>
                <Table className={"records-table"}>
                    <TableBody>
                        {this.renderRecords(recordTypes, records)}
                    </TableBody>
                </Table>
            </TableContainer>
        )
    }
}
