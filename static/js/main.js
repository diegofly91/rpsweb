
document.querySelector('.btn').addEventListener('click', function() {
});

let imgArray = [
    "static/img/hand_rock.png",
    "static/img/hand_paper.png",
    "static/img/hand_scissors.png",
]
let chooseArray = [
    "PIEDRA",
    "PAPEL",
    "TIJERA",
] 

function choose(x) {
    fetch('/play?c=' + x)
    .then(response => response.json())
    .then(data => {
        // 
        document.getElementById("player_choose").innerHTML = 'El jugador a elejido <strong>'+chooseArray[x] +'</strong>';
        document.getElementById("computer_choose").innerHTML = 'la Computadora a elejido <strong>'+chooseArray[data.computer_choice_int] +'</strong>';
        document.getElementById("player_score").innerHTML = data.player_score;

        document.getElementById("img_computer").setAttribute('src', imgArray[data.computer_choice_int]);
        document.getElementById("computer_score").innerHTML = data.computer_score;
        document.getElementById("round_message").innerHTML = data.message;
        document.getElementById('round_result').innerHTML = data.round_result;
    });
    

}