function initComm() {

    function error_handler(msg) {
        console.log("Recv'd ERROR: " + msg)
    }

    function ack_handler(msg) {
        console.log("Recv'd ACK: " + msg)
    }

    function filecmd_handler(msg) {
        console.log("Recv'd FILECMD: " + msg)
    }

    $("#send_msg_btn").click(function() {
        websock_comm.Send($('#comm_input').val());
    })

    websock_comm.AddHandler("error", error_handler)
    websock_comm.AddHandler("ack", ack_handler)
    websock_comm.AddHandler("filecmd_response", filecmd_handler)

    websock_comm.Connect(window.location.host, "/comm");
}