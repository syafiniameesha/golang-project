import { ElNotification } from 'element-plus';

export function notification(title, message, position) {
    if (position == null || position == undefined || position == "") {
        position = "bottom-right"
    }
    ElNotification({
        title: title,
        message: message,
        position: position,
        icon: 'Check',
    });
}


