import React from 'react'
import RecordsMenu from "./RecordsMenu";
import RecordsReader from "./RecordsReader";

export default class RecordsWrapper extends React.Component {
    constructor(props) {
        super(props);
        this.records = props.records
    }

    render() {
        return (
            <div>
                <RecordsMenu records={this.records}/>
                <RecordsReader/>
            </div>
        )
    }
}
