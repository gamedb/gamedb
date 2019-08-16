if ($('#home-page').length > 0) {

    // Players
    $('[data-sort]').on('click', function (e) {

        loadPlayers($(this).attr('data-sort'));
    });

    loadPlayers('level');

    function loadPlayers(sort) {

        $('#players .fa-spin').removeClass('d-none');
        $('#players table').addClass('d-none');

        $('[data-sort]').removeClass('badge-success');
        $('[data-sort="' + sort + '"]').addClass('badge-success');

        $('#column').html(sort);

        $.ajax({
            url: '/home/' + sort + '/players.json',
            dataType: 'json',
            cache: true,
            success: function (data, textStatus, jqXHR) {

                // Reset, for when changing order
                $('#players .fa-spin').addClass('d-none');
                $('#players table').removeClass('d-none');
                $('#players tbody tr').remove();

                if (isIterable(data)) {

                    const $container = $('#players tbody');

                    $container.json2html(
                        data,
                        {
                            '<>': 'tr', 'data-link': '${link}', 'html': [
                                {
                                    '<>': 'td', 'class': 'font-weight-bold', 'html': '${rank}'
                                },
                                {
                                    '<>': 'td', 'class': 'img', 'html': [
                                        {
                                            '<>': 'div', 'class': 'icon-name', 'html': [
                                                {
                                                    '<>': 'div', 'class': 'icon', 'html': [{'<>': 'img', 'data-lazy': '${avatar}'}],
                                                },
                                                {
                                                    '<>': 'div', 'class': 'name', 'html': '${name}'
                                                }
                                            ]
                                        }
                                    ]
                                },
                                {
                                    '<>': 'td', 'nowrap': 'nowrap', 'class': function () {
                                        if (sort === 'level') {
                                            return 'img';
                                        } else {
                                            return '';
                                        }
                                    }, 'html': function () {

                                        switch (sort) {
                                            case 'level':
                                                return '<div class="icon-name"><div class="icon"><div class="' + this.class + '"></div></div><div class="name min">' + this.value + '</div></div>';
                                            case 'games':
                                                return this.value + ' games';
                                            case 'badges':
                                                return this.value + ' badges';
                                            default:
                                                return this.value;
                                        }
                                    },
                                },
                            ]
                        },
                        {
                            prepend: false,
                        }
                    );

                    observeLazyImages('#players img[data-lazy]');
                }
            },
        });
    }

    // Prices
    $.ajax({
        url: '/home/prices.json',
        dataType: 'json',
        cache: false,
        success: function (data, textStatus, jqXHR) {

            $('#prices .fa-spin').remove();
            $('#prices table').removeClass('d-none');

            addPriceRow(data, false);
        },
    });

    websocketListener('prices', function (e) {

        const data = $.parseJSON(e.data);

        if (data.Data[13] === user.prodCC) { // CC
            if (data.Data[12] < 0) { // Drops
                if (data.Data[0] > 0) { // Apps
                    addPriceRow([{
                        "name": data.Data[3],
                        "id": data.Data[0],
                        "link": data.Data[5],
                        "after": data.Data[7],
                        "discount": data.Data[14],
                        "time": data.Data[11],
                        "avatar": data.Data[4],
                    }], true);
                }
            }
        }
    });

    function addPriceRow(data, addToTop) {

        const $container = $('#prices tbody');

        $container.json2html(
            data,
            {
                '<>': 'tr', 'data-link': '${link}', 'html': [
                    {
                        '<>': 'td', 'class': 'img', 'html': [
                            {
                                '<>': 'div', 'class': 'icon-name', 'html': [
                                    {
                                        '<>': 'div', 'class': 'icon', 'html': [{'<>': 'img', 'data-lazy': '${avatar}'}]
                                    },
                                    {
                                        '<>': 'div', 'class': 'name', 'html': '${name}'
                                    }
                                ],
                            },
                        ]
                    },
                    {
                        '<>': 'td', 'html': '${after}', 'nowrap': 'nowrap'
                    },
                    {
                        '<>': 'td', 'html': '${discount}%', 'nowrap': 'nowrap'
                    },
                    {
                        '<>': 'td', 'nowrap': 'nowrap', 'html': [
                            {
                                '<>': 'span', 'data-toggle': 'tooltip', 'data-placement': 'left', 'data-livestamp': '${time}',
                            }
                        ],
                    },
                ]
            },
            {
                prepend: addToTop,
            }
        );

        $container.find('tr').slice(15).remove();

        observeLazyImages('#prices img[data-lazy]');
    }
}
