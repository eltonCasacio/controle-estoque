create table usuarios(
    id varchar(255),
    nome varchar(255),
    senha varchar(255),
    ativo bool,
    PRIMARY KEY (id)
);

create table fornecedores(
    id varchar(255),
    razao_social varchar(255),
    nome_fantasia varchar(255),
    cnpj varchar(20),
    ie varchar(30),
    ativo bool,
    PRIMARY KEY (id)
);

create table fornecedores_pecas(
    fornecedor_id varchar(255),
    id_peca varchar(255)
);

create table enderecos(
    id int unsigned not null auto_increment,
    fornecedor_id varchar(255),
    uf varchar(255), 
    rua varchar(255), 
    complemento varchar(255), 
    bairro varchar(255), 
    cep varchar(15), 
    numero varchar(10), 
    PRIMARY KEY (id),
    CONSTRAINT fk_enderecos_fornecedores FOREIGN KEY (fornecedor_id) REFERENCES fornecedores(id)
);

create table contatos(
    id int unsigned not null auto_increment,
    fornecedor_id varchar(255),
    email varchar(100), 
    celular varchar(20), 
    nome varchar(255), 
    PRIMARY KEY (id),
    CONSTRAINT fk_contatos_fornecedores FOREIGN KEY (fornecedor_id) REFERENCES fornecedores(id)
);

create table pecas(
    id varchar(255), 
    fornecedor_id varchar(255), 
    codigo varchar(255), 
    descricao varchar(255), 
    materiaprima varchar(255), 
    url_desenho_tecnico varchar(255), 
    url_foto varchar(255), 
    descricao_tecnica varchar(255),
    massa varchar(255), 
    quantidade varchar(255),
    PRIMARY KEY (id),
    CONSTRAINT fk_pecas_fornecedores FOREIGN KEY (fornecedor_id) REFERENCES fornecedores(id)
);
