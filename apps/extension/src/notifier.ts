import * as vscode from 'vscode';

export function registerNotificationTrigger(context: vscode.ExtensionContext) {
    const disposable = vscode.commands.registerCommand('devhub.sendDiscordNotification', async () => {
        const title = await vscode.window.showInputBox({ prompt: 'Enter Notification Title' });
        if (!title) return;

        const message = await vscode.window.showInputBox({ prompt: 'Enter Notification Message' });
        if (!message) return;

        try {
            const config = vscode.workspace.getConfiguration('devhub');
            const backendUrl = config.get<string>('backendUrl') || 'http://localhost:8080';

            const res = await fetch(`${backendUrl}/api/v1/discord/notify`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    event_type: 'task_update',
                    title: title,
                    message: message
                })
            });

            if (res.ok) {
                vscode.window.showInformationMessage('Discord Notification dispatched successfully!');
            } else {
                vscode.window.showErrorMessage('Failed to trigger Discord notification');
            }
        } catch (err: any) {
            vscode.window.showErrorMessage(`Error triggering notification: ${err.message}`);
        }
    });

    context.subscriptions.push(disposable);
}
