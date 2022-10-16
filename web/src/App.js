import React from "react"
import {Card, CardContent, Container, TextField} from '@mui/material'
import {ThemeProvider, createTheme} from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import RecordsWrapper from "./component/RecordsWrapper"

const darkTheme = createTheme({
    palette: {
        mode: 'dark'
    }
})


export default class App extends React.Component {
    constructor(props) {
        super(props)
        this.state = {hostname: ""}
        this.hostInputRef = React.createRef()
    }

    changeHostname(e) {
        if (e.key === 'Enter') {
            this.setState({hostname: this.hostInputRef.current.value})
        }
    }

    render() {
        return (
            <ThemeProvider theme={darkTheme}>
                <CssBaseline/>
                <Container className={"container"}>
                    <Card elevation={1}>
                        <CardContent>
                            <TextField
                                fullWidth
                                required
                                inputRef={this.hostInputRef}
                                onKeyDown={this.changeHostname.bind(this)}
                                label="Hostname"
                                variant="outlined"/>
                            <RecordsWrapper hostname={this.state.hostname}
                                            records={this.props.records}></RecordsWrapper>
                        </CardContent>
                    </Card>
                </Container>
            </ThemeProvider>
        )
    }
}
