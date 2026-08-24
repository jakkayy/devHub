import * as http from 'http';
import * as vscode from 'vscode';

export class LocalMockServer {
    private _server?: http.Server;
    private _port: number = 9090;

    public startServer() {
        if (this._server) {
            return;
        }

        this._server = http.createServer((req, res) => {
            res.setHeader('Access-Control-Allow-Origin', '*');
            res.setHeader('Content-Type', 'application/json');

            if (req.url === '/mock/api/v1/tasks') {
                res.writeHead(200);
                res.end(JSON.stringify({
                    status: 'mock_success',
                    data: [
                        { id: 'MOCK-1', title: 'Local Mock Task 1', assignee: 'IDE User', status: 'In Progress' }
                    ]
                }));
                return;
            }

            res.writeHead(200);
            res.end(JSON.stringify({
                message: 'Local devHub Mock Server is Active',
                path: req.url,
                timestamp: new Date().toISOString()
            }));
        });

        this._server.listen(this._port, () => {
            console.log(`devHub Local Mock Server running at http://localhost:${this._port}`);
        });
    }

    public stopServer() {
        if (this._server) {
            this._server.close();
            this._server = undefined;
        }
    }
}
