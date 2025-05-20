const WS_URL_LIVE = "https://ws.inshasaki.com/api";
const WS_TOKEN_LIVE = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOjExMjE2LCJpc3MiOiJodHRwczovL3dzLmluc2hhc2FraS5jb20vYXBpL2F1dGgvbG9naW4iLCJpYXQiOjE3MzI4NTg4NzQsImV4cCI6MTc1OTEzODg3NCwibmJmIjoxNzMyODU4ODc0LCJqdGkiOiJKb3hnZThQZ0dwRm54a25lIn0.prsgoL4rQT2X-G49OmCfOcMBzXpB9bwF6OJpjTmev2E"

const WS_URL_QC = "https://test-ws.inshasaki.com/api";
const WS_TOKEN_QC = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOjE3MjYxLCJpc3MiOiJodHRwczovL3Rlc3Qtd3MuaW5zaGFzYWtpLmNvbS9hcGkvYXV0aC9sb2dpbiIsImlhdCI6MTcyODM2MjE0OSwiZXhwIjoxNzU0NjQyMTQ5LCJuYmYiOjE3MjgzNjIxNDksImp0aSI6InhOdmRaMFRraTVjd0U0Z3MifQ.rGoSpuFgPAEqJ5MnDs2Rte2cFVzuhj1XivpxFRARlD4"

const getListReceipt = async (params) => {
    let url = WS_URL_LIVE + "/report/receipt";

    if (params) {
        url += "?";
        Object.keys(params).forEach((key, index) => {
            url += key + "=" + params[key];
            if (index < Object.keys(params).length - 1) {
                url += "&";
            }
        });
    }

    let token = WS_TOKEN_LIVE;
    let response = await fetch(url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
    });
    let data = await response.json();
    return data.data.rows;
}

const fixReceiptDebit = async (receipt_id) => {
    let url = WS_URL_LIVE + "/accounting/receipt/" + receipt_id;

    let token = WS_TOKEN_LIVE;
    let response = await fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
        body: JSON.stringify({
            "receipt_balance": 0,
            "is_optimize": 1
        })
    });
    let data = await response.json();
    return data;
}

const checkPaymentAmout = (data) => {
    const receipt = data ?? null;
    const payments = data?.payments ?? [];
    if (!receipt || payments.length == 0) {
        return false;
    }

    if (receipt.receipt_balance == 0) {
        return "DONE " + receipt.receipt_code;
    }

    const totalPayment = payments.reduce((total, payment) => {
        return total + payment.payment_amount;
    }, 0);

    if (totalPayment == receipt.receipt_subtotal) {
        console.log("Receipt OK: ", receipt.receipt_id, receipt.receipt_code, totalPayment, receipt.receipt_subtotal, receipt.receipt_balance);
        fixReceiptDebit(receipt.receipt_id);
        return {
            receipt: receipt,
            totalPayment: totalPayment
        };
    } else {
        console.log({
            receipt: receipt,
            totalPayment: totalPayment
        })
        return false;
    }
}

const formatDate = (date) => {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
}

const main = async () => {
    document.getElementById("loading").style.display = "block";
    const dateTime = document.getElementById("datetime");
    const toDate = new Date(dateTime.value);

    const from = document.getElementById("fromDate");
    const fromDate = new Date(from.value);
    const params = {
        from_date: formatDate(fromDate),
        to_date: formatDate(toDate),
        status: 2,
        limit: 100,
        receipt_balance: 1000,
        without_ext: 1,
        config: 256,
        join_payments: 1
    }

    try {

        const listReceipt = await getListReceipt(params);

        const checkPaymentPromises = listReceipt.map((receipt) => checkPaymentAmout(receipt));

        const results = await Promise.all(checkPaymentPromises);
        const arrRs = results.filter((rs) => rs != false);
        let viewer = document.getElementById("viewer");

        let html = "<table id='table'> <tr> <th>No</th> <th>Receipt ID</th> <th>Receipt Code</th> <th>Subtotal</th> <th>Total Payment</th> <th>Balance</th> </tr>";
        arrRs.forEach((rs, index) => {
            let tr = "<tr>";
            tr += "<td class='text-center'>" + (index + 1) + "</td>";
            tr += "<td class='text-center'>" + rs.receipt.receipt_id + "</td>";
            tr += "<td class='text-center'>" + rs.receipt.receipt_code + "</td>";
            tr += "<td class='text-center'>" + rs.receipt.receipt_subtotal + "</td>";
            tr += "<td class='text-center'>" + rs.totalPayment + "</td>";
            tr += "<td class='text-center'>" + rs.receipt.receipt_balance + "</td>";
            tr += "</tr>";
            html += tr;
        });
        html += "</table>";
        console.log(html)
        document.getElementById("loading").style.display = "none";
        viewer.innerHTML = html;
    } catch (error) {
        document.getElementById("loading").style.display = "none";
        console.error("Error in main function:", error);
    }
}

