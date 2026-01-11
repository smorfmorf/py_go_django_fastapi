// import { exec } from "child_process";

// exec("wmic bios get serialnumber", (error, stdout, stderr) => {
//     if (error) {
//         console.error(`Ошибка: ${error.message}`);
//         return;
//     }
//     if (stderr) {
//         console.error(`stderr: ${stderr}`);
//         return;
//     }

//     // stdout содержит две строки: "SerialNumber" и сам номер
//     const lines = stdout.trim().split("\n");
//     const serialNumber = lines[1].trim(); // вторая строка — серийник
//     console.log("Serial Number:", serialNumber);
// });

// import crypto from 'crypto';



const crypto = require('crypto');

// const hash = crypto.pbkdf2('password', 'salt', 10_000_000, 64, 'sha512', (err, res) => {
//     console.log('Hashing...', res.toString('hex'));
// });
// const hash2 = crypto.pbkdf2('password', 'salt', 10_000_000, 64, 'sha512', (err, res) => {
//     console.log('Hashing2...', res.toString('hex'));
// });

// 👉 Event loop заблокирован, пока крутится for

function doWork() {
    return new Promise((resolve) => {

        console.log('Start long loop...');
        resolve(123);

        const hash = crypto.pbkdf2('password', 'salt', 10_000_000, 64, 'sha512', (err, res) => {
            console.log('Hashing...', res.toString('hex'));
        });
        // for (let i = 0; i < 1e9; i++) {
        //     for (let i = 0; i < 4; i++) {
        //         // long loop
        //     }
        // }
        console.log('End long loop...');

    });
}

async function main() {
    doWork().then((res) => {
        console.log('doWork result:', res);
    });
    console.log('after doWork');
}

main();




















const http = require('http')
const express = require('express')
const WebSocket = require('ws')

const app = express()
const server = http.createServer(app)

const wss = new WebSocket.Server({ server })

wss.on('connection', function connection(ws) {
    ws.id = 'Client-' + Math.random().toString(16).slice(2);

    ws.on('error', console.error);

    ws.on('message', function message(data) {
        console.log(`received: ${ws.id}`, data.toString());

        // отправляем всем подключённым клиентам (в том числе отправителю)
        wss.clients.forEach(client => {
            if (client.readyState === WebSocket.OPEN && client !== ws) {
                client.send(`${ws.id} says: ${data.toString()}`);
            }
        });
    });
    console.log('Client connected', ws.id);
    ws.send('something');
});


server.listen(3006, () => {
    console.log('Server started on: http://localhost:3006');
})
