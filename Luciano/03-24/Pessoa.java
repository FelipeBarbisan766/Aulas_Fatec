public class Pessoa{
    private String nome;
    private String cpf;
    private String email;
    private String dtnasc;
    
    public Pessoa(String nome,String cpf,String email, String dtnasc){
        this.nome = nome;
        this.cpf = cpf;
        this.email = email;
        this.dtnasc = dtnasc;
    }
    public void mostrar(){
        System.out.println("Nome: "+ nome +"\n"+
        "CPF: "+ cpf +"\n"+
        "E-mail: "+ email +"\n"+
        "Data de Nascimento: "+ dtnasc);
    }
}