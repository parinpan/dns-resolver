import React from 'react'
import RecordsMenu from "./RecordsMenu";
import RecordsReader from "./RecordsReader";

export default class RecordsWrapper extends React.Component {
    constructor(props) {
        super(props)
        this.state = {queryRecord: "ANY"}
    }

    changeRecord(record) {
        this.setState({queryRecord: record})
    }

    render() {
        return (
            <div>
                <RecordsMenu changeRecord={this.changeRecord.bind(this)} records={this.props.records}/>
                <RecordsReader queryHostname={this.props.hostname} queryRecord={this.state.queryRecord}/>
            </div>
        )
    }
}
