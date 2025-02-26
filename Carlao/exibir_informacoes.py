class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano

    def exibir_informacoes(self):
        print(f"Marca: {self.marca} Modelo: {self.modelo} Ano: {self.ano}")
        
        
# Instanciando um objeto da classe Carro
fusca = Carro("Volkswagen", "Fusca", 1970)
fusca.exibir_informacoes()