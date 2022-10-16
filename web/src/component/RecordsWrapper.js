import React from "react"
import RecordsMenu from "./RecordsMenu";
import RecordsReader from "./RecordsReader";
import {Accordion, AccordionDetails, AccordionSummary, Typography} from "@mui/material";
import RecordFetcher from "../api/RecordFetcher";
import Placeholder from "./Placeholder";

export default class RecordsWrapper extends React.Component {
    constructor(props) {
        super(props)
        this.fetcher = new RecordFetcher('http://localhost')

        this.state = {
            results: {},
            nsResults: {},
            queryRecord: "ANY",
        }
    }

    componentDidUpdate(prevProps) {
        if (prevProps.hostname !== this.props.hostname) {
            this.fetchData(this.props.hostname, this.state.queryRecord)
        }
    }

    fetchData(hostname, record) {
        const self = this

        if (!this.validState(hostname, record)) {
            return
        }

        this.fetcher.fetch(hostname, record)
            .then(response => {
                self.setState({results: response.data})
            })
            .catch(e => {
                console.log(e)
            })

        this.fetcher.fetchNS(hostname, record)
            .then(response => {
                self.setState({nsResults: response.data})
            })
            .catch(e => {
                console.log(e)
            })
    }

    changeRecord(record) {
        this.setState({queryRecord: record})
        this.fetchData(this.props.hostname, record)
    }

    renderResults(hostname, results) {
        const notEmpty = Object.keys(results || {}).length > 0

        return (
            <Accordion className={"accordion"} key={hostname} expanded={true} elevation={5}>
                <AccordionSummary>
                    <Typography>Answer from {hostname}</Typography>
                </AccordionSummary>
                <AccordionDetails>
                    {this.renderSwitch(notEmpty,
                        (<RecordsReader records={results}/>),
                        (<Placeholder text={"Records not found!"}/>))}
                </AccordionDetails>
            </Accordion>
        )
    }

    mergeRenderResults(hostname, results, nsResults) {
        const merged = []
        const nameservers = Object.keys(nsResults)
        merged.push(this.renderResults(hostname, results))

        nameservers.forEach(ns => {
            const title = `${ns} (NS of ${hostname})`
            merged.push(this.renderResults(title, nsResults[ns]))
        })

        return merged
    }

    renderPlaceholder(text) {
        return (<Placeholder text={text}/>)
    }

    validState(hostname, record) {
        return (hostname || "") !== "" && (record || "") !== ""
    }

    renderSwitch(condition, left, right) {
        return condition ? left : right
    }

    render() {
        const {hostname} = this.props
        const {queryRecord} = this.state

        return (
            <div className={"records-wrapper"}>
                <RecordsMenu changeRecord={this.changeRecord.bind(this)} records={this.props.records}/>
                {this.renderSwitch(
                    this.validState(hostname, queryRecord),
                    this.mergeRenderResults(hostname, this.state.results, this.state.nsResults),
                    this.renderPlaceholder("Hostname and record type must be specified"))}
            </div>
        )
    }
}
