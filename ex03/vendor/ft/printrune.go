/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printrune.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 16:13:38 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/25 16:14:11 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package ft

import (
	"os"
	"errors"
	"unicode/utf8"
)

func PrintRune(r rune) error{
	l := utf8.RuneLen(r)
	if l == -1 {
		return errors.New("The rune is not a valid value to encode in UTF-8")
	}
	p := make([]byte, l)
	utf8.EncodeRune(p, r)
	_, err := os.Stdout.Write(p)
	return err
}
