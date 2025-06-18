$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);
$('#editar-usuario').on('submit', editar);
$('#atualizar-senha').on('submit', atualizarSenha);
$('#deletar-usuario').on('click', deletarUsuario);

function pararDeSeguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: 'POST'
    }).done(function() {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário", "error");
        $('#parar-de-seguir').prop('disabled', false);
    });

}

function seguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: 'POST'
    }).done(function() {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao seguir o usuário", "error");
        $('#seguir').prop('disabled', false);
    });
}

function editar(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success").then(function() {
            window.location = "/perfil";
        });
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error");
    });
}

function atualizarSenha(evento) {
    evento.preventDefault();
    var senhaAtual = $('#senha-atual').val();
    var novaSenha = $('#nova-senha').val();
    var confirmarSenha = $('#confirmar-senha').val();

    if (novaSenha !== confirmarSenha) {
        Swal.fire("Ops...", "As senhas não coincidem!", "warning");
        return;
    }

    $.ajax({
        url: "/atualizar-senha",
        method: "POST",
        data: {
            atual: senhaAtual,
            nova: novaSenha
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Senha atualizada com sucesso!", "success").then(function () {
            window.location = "/perfil";
        });
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar a senha!", "error");
    });

}

function deletarUsuario(evento) {
    evento.preventDefault();
    console.log("deletar usuario");
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir sua conta? Essa ação é irreversível!",
        confirmButtonText: "Sim, excluir",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning",
    }).then((confirmacao) => {
        console.log("THEN");
        if (!confirmacao.value) return; 
        
        $.ajax({
            url: "/deletar-usuario",
            method: "DELETE"
        }).done(function() {
            Swal.fire("Sucesso!", "Conta deletada com sucesso!", "success").then(function() {
                window.location = "/logout";
            });
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao deletar a conta!", "error");
        });
        
    });
}
