import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';

const root = ReactDOM.createRoot(
    document.getElementById('root')
);

const records = [
    "ANY", "A", "AAAA", "CAA", "CNAME", "DNSKEY", "DS", "MX",
    "NS", "PTR", "SOA", "SRV", "SOA", "SRV", "TLSA", "TSIG", "TXT"
]

root.render(
    <React.StrictMode>
        <App records={records}/>
    </React.StrictMode>
);
