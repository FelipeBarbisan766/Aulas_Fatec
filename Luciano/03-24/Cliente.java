public class Cliente extends Pessoa {
    private String rg;
    
    public Cliente(String rg, String nome, String cpf, String email, String dtnasc){
        super(nome,cpf,email,dtnasc);
        this.rg = rg;
    }
    @Override
    public String toString(){
        mostrar();
        return "RG: "+ rg ;
        
    }
}