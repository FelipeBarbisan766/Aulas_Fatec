// Exemplo de função, pois há return arquivofunção.html
function Principal() {
    var base, altura, reultado;
    base = Number(window.prompt("Digite Base"));
    altura = Number(window.prompt("Digite altura"));
    reultado = calc_area(base, altura);
    document.write(reultado);
}
function calc_area(p1, p2) {
    var res;
    res = p1 * p2;
    return res;
}
// Exemplo procedimento pois não ha return arquivoprocedimento.html
// function Principal(){
//     var base,altura,reultado;
//     base =  Number(window.prompt("Digite Base"));
//     altura =  Number(window.prompt("Digite altura"));
//     calc_area(base,altura);

// }
// function calc_area(p1,p2){
//     var res; 
//     res = p1 * p2;
//     document.write(res);
// }


function fun_ex1_pri() {
    var v1;
    v1 = Number(window.prompt("Digite um valor"));
    fun_ex1_calc(v1);
}
function fun_ex1_calc(p1) {
    if (p1 < 100) {
        document.write("Numero menor que 100");
    } else {
        document.write("Numero igual ou maior que 100");
    }
}

function fun_ex2_pri() {
    var v1, v2, reultado;
    v1 = Number(window.prompt("Digite um valor"));
    v2 = Number(window.prompt("Digite outro valor"));
    reultado = fun_ex2_calc(v1, v2);
    document.write(reultado);
}
function fun_ex2_calc(p1, p2) {
    var res;
    res = p1 + p2;
    return res;
}

function fun_ex3_pri() {
    var v1;
    v1 = Number(window.prompt("Digite um numero"));
    fun_ex3_calc(v1);
}
function fun_ex3_calc(p1) {
    switch (p1) {
        case 1:
            document.write("Secretaria");
            break;
        case 2:
            document.write("Gerente");
            break;
        case 3:
            document.write("Operario");
            break;
        default:
            document.write("outros");
            break;
    }
}

function fun_ex4_pri() {
    var v1;
    v1 = Number(window.prompt("Digite um numero"));
    fun_ex4_calc(v1);
}
function fun_ex4_calc(p1) {
    var k,m,g;
    if(p1 >= 1024){
        k = p1/1024;
    }
    if(k >= 1024){
        m = k/1024;
    }
    if(m >= 1024){
        g = m/1024;
    }
    document.write(p1+" Bytes é Igual a "+k+" kilobytes , "+m+" Megabytes , "+g+" Gigabytes")
}