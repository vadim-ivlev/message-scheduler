var urlParams = new URLSearchParams(window.location.search)
var messageSchedulerURL = urlParams.get('end_point') ? urlParams.get('end_point') : 'http://localhost:8088/'
var functionUrl = urlParams.get('function_url') ? urlParams.get('function_url') : 'https://us-central1-rg-push.cloudfunctions.net/'

// if (document.location.hostname == 'localhost') {
//     functionUrl = 'http://localhost:5001/rg-push/us-central1/'
// }  

function subscribeTokenToTopic(token,topic) {
    fetch( functionUrl +`subscribe_token_to_topic?iid=${token}&topic=${topic}`)
        .then(res => res.json())
        .then(json => console.log(json))
        .catch(err => console.log("ERROR:",err))
}

function unsubscribeTokenFromTopic(token, topic) {
    fetch(functionUrl +`unsubscribe_token_from_topic?iid=${token}&topic=${topic}`)
        .then(res => res.json())
        .then(json => console.log(json))
        .catch(err => console.log("ERROR:",err))
}


function createMessageFirebase(to, message, link, wait, /*status,*/ user){
    firebase.database().ref('/messages').push({to, message, link, wait, /*status,*/ user})
}

function createMessage(to, message, link, wait, /*status,*/ user){

    let query =`
    mutation {
        create_message(
            message: "${message}",
            link: "${link}",
            wait: ${wait},
            to: "${to}"
        )
    }
    `

    fetch(messageSchedulerURL+`schema`, { method: 'POST', credentials: 'include', body: JSON.stringify({ query: query, variables: {} }) })
        .then(res => res.json())
        .then((json) => {
            console.log(json)
            json.errors && console.log(json.errors[0].message)
        })
        .catch(err => console.log("ERROR:",err))  
}


function deleteMessageFirebase(key){
    firebase.database().ref('/messages').child(key).remove()
}

async function editMessage(key){
    var ref = firebase.database().ref('/messages').child(key)
    var snap = await ref.once('value')
    var msg = snap.val()
    document.getElementById('to').value = msg.to 
    document.getElementById('txt').value = msg.message
    document.getElementById('link').value = msg.link
    document.getElementById('wait').value = msg.wait
    console.log(msg)
    window.scrollTo(0,0)
    document.getElementById('txt').focus()
    document.getElementById('main-area').style.backgroundColor = 'silver'
    setTimeout(()=>{ document.getElementById('main-area').style.backgroundColor = 'whitesmoke' },100)
    // ref.remove()
}


function sendScheduledMessages(){
    fetch(functionUrl +`send_scheduled_messages`)
        .then(res => res.json())
        .then(json => console.log(json))
        .catch(err => console.log("sendScheduledMessages ERROR:",err))
}

