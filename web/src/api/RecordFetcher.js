export default class RecordFetcher {
    constructor(host) {
        this.host = host
    }

    async fetch(hostname, record) {
        const response = await fetch(this.host + '/resolve', {
            method: 'POST',
            redirect: 'follow',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                'hostname': hostname,
                'record_type': record
            })
        })

        return await response.json()
    }


    async fetchNS(hostname, record) {
        const response = await fetch(this.host + '/resolve/ns', {
            method: 'POST',
            redirect: 'follow',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                'hostname': hostname,
                'record_type': record
            })
        })

        return await response.json()
    }
}
