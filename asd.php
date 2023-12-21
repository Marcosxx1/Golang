interface Amor {
    public function paixão();
}

class Marcos implements Amor {
    public function paixão() {
        return "Amo você amorzão";
    }
}

class Aline implements Amor {
    public function paixão() {
        return "Amo você amorzinho";
    }
}

$marcos = new Marcos();
$aline = new Aline();

$MarcosTeAma = $marcos->paixão();
$FofuraMeAma = $aline->paixão();

echo $MarcosTeAma;
echo $FofuraMeAma;
