$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
  evento.preventDefault();

  if ($('#senha').val() != $('#confirmar-senha').val()) {
    alert('As senhas não coincidem!');
    return;
  }

  var resposta = $.ajax({
    url: '/usuarios',
    method: 'POST',
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      senha: $('#senha').val(),
    },
  });
  resposta.done(function () {
    alert('Usuário cadastrado com sucesso!');
  });
  resposta.fail(function () {
    alert('Erro ao cadastrar usuário');
  });
}
