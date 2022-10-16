import React from 'react'
import {Card, CardContent, Container, Grid, TextField} from '@mui/material'
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
    }

    render() {
        return (
            <ThemeProvider theme={darkTheme}>
                <CssBaseline/>
                <Container>
                    <Card elevation={1}>
                        <CardContent>
                            <TextField fullWidth required id="outlined-basic" label="Hostname" variant="outlined"/>
                            <RecordsWrapper records={this.props.records}></RecordsWrapper>
                        </CardContent>
                    </Card>
                </Container>
            </ThemeProvider>
        )
    }
}
