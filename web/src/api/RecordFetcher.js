export default class RecordFetcher {
    constructor(host) {
        this.endpoint = host + '/resolve'
    }

    async fetch(hostname, record) {
        const response = await fetch(this.endpoint, {
            method: 'POST',
            redirect: 'follow',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                'hostname': hostname,
                'record_type': record
            })
        })

        return await response.json()
    }
}
