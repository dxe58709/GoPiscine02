/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sqrt.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/27 16:31:20 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/27 16:37:02 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	for i := 0; i * i <= nb; i++ {
		if i * i == nb {
			return i
		}
	}
	return 0
}