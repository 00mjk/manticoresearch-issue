<?php

   require_once __DIR__ . '/vendor/autoload.php';

   $config = ['host'=>'127.0.0.1','port'=>9308];
   $client = new \Manticoresearch\Client($config);
   $params = [
        'body' => [
            'index' => 'rt_pneus_illico_catalog2',
            'query' => [
                'match_phrase' => [
                    'manufacturer' => 'Fulda',
                ]
            ]
        ]
    ];
    $response = $client->search($params);
    print_r($response)

?>

